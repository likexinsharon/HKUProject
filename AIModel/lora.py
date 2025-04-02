from datasets import Dataset
import pandas as pd
from transformers import (
    AutoTokenizer,
    AutoModelForCausalLM,
    DataCollatorForSeq2Seq,
    TrainingArguments,
    GenerationConfig,
    Trainer, )
import torch, os
from peft import LoraConfig, TaskType, get_peft_model
import warnings


# 对数据集进行处理，需要将数据集的内容按大模型的对话格式进行处理
def process_func(example):
    MAX_LENGTH = 384    # Llama分词器会将一个中文字切分为多个token，因此需要放开一些最大长度，保证数据的完整性
    input_ids, attention_mask, labels = [], [], []
    prompt = ''
    instruction = tokenizer(f"User: {prompt}{example['question']}\n\n", add_special_tokens=False)
    thought_process = example['thought_process'][0] if example['thought_process'] else ""
    full_answer = f"{thought_process} {example['answer']}"
    response = tokenizer(f"Assistant: {full_answer}<｜end▁of▁sentence｜>", add_special_tokens=False)
    # response = tokenizer(f"Assistant: {example['output']}<｜end▁of▁sentence｜>", add_special_tokens=False)
    input_ids = instruction["input_ids"] + response["input_ids"] + [tokenizer.pad_token_id]
    attention_mask = instruction["attention_mask"] + response["attention_mask"] + [1]  # 因为eos token咱们也是要关注的所以 补充为1
    labels = [-100] * len(instruction["input_ids"]) + response["input_ids"] + [tokenizer.pad_token_id]
    if len(input_ids) > MAX_LENGTH:  # 做一个截断
        input_ids = input_ids[:MAX_LENGTH]
        attention_mask = attention_mask[:MAX_LENGTH]
        labels = labels[:MAX_LENGTH]
    return {
        "input_ids": input_ids,
        "attention_mask": attention_mask,
        "labels": labels
    }
if __name__ == "__main__":
    device = 'cuda' if torch.cuda.is_available() else 'cpu'
    # 模型文件路径
    model_path = ''
    # 训练过程数据保存路径
    output_dir = ''
    # 是否从上次断点处接着训练
    train_with_checkpoint = False

    # 加载数据集
    df = pd.read_json('')
    df['thought_process'] = df['thought_process'].apply(lambda x: x[0] if x else "")
    ds = Dataset.from_pandas(df)

    # 加载模型
    model = AutoModelForCausalLM.from_pretrained(model_path, device_map=device, torch_dtype=torch.bfloat16, use_cache=False)
    model.generation_config = GenerationConfig.from_pretrained(model_path)
    model.generation_config.pad_token_id = model.generation_config.eos_token_id
    print(model)
    model.enable_input_require_grads()  # 开启梯度检查点时，要执行该方法

    # 加载tokenizer
    tokenizer = AutoTokenizer.from_pretrained(model_path, use_fast=False, trust_remote_code=True)
    tokenizer.padding_side = 'right'
        # 应用 process_func 到数据集
    tokenized_ds = ds.map(process_func, batched=False)

    config = LoraConfig(
        task_type=TaskType.CAUSAL_LM,
        target_modules=["q_proj", "k_proj", "v_proj", "o_proj", "gate_proj", "up_proj", "down_proj"],
        inference_mode=False,  # 训练模式
        r=8,  # Lora 秩
        lora_alpha=32,  # Lora alaph，具体作用参见 Lora 原理
        lora_dropout=0.1  # Dropout 比例
    )

    model = get_peft_model(model, config)
    model.print_trainable_parameters()
    args = TrainingArguments(
        output_dir=output_dir,
        per_device_train_batch_size=2,
        gradient_accumulation_steps=2,
        logging_steps=20,
        num_train_epochs=2,
        save_steps=25,
        save_total_limit=2,
        learning_rate=1e-4,
        save_on_each_node=True,
        gradient_checkpointing=True
    )
    trainer = Trainer(
        model=model,
        args=args,
        train_dataset=tokenized_ds,
        data_collator=DataCollatorForSeq2Seq(tokenizer=tokenizer, padding=True),
    )
    # 如果训练中断了，还可以从上次中断保存的位置继续开始训练
    if train_with_checkpoint:
        checkpoint = [file for file in os.listdir(output_dir) if 'checkpoint' in file][-1]
        last_checkpoint = f'{output_dir}/{checkpoint}'
        print(last_checkpoint)
        trainer.train(resume_from_checkpoint=last_checkpoint)
    else:
        trainer.train()
    text = "我家的猫有猫鼻支，但是为什么平时不会打喷嚏，但是只有冬天会"
    inputs = tokenizer(f"User: {text}\n\n", return_tensors="pt")
    outputs = model.generate(**inputs.to(model.device), max_new_tokens=100)

    result = tokenizer.decode(outputs[0], skip_special_tokens=True)
    print(result)