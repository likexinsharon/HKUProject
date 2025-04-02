from transformers import AutoModelForCausalLM, AutoTokenizer
import torch
from peft import PeftModel
from peft import LoraConfig, TaskType, get_peft_model
 
model_path = ''
lora_path = ""
device = 'cuda' if torch.cuda.is_available() else 'cpu'
# 合并后的模型路径
output_path = r''
 
# 等于训练时的config参数
config = LoraConfig(
    task_type=TaskType.CAUSAL_LM,
    target_modules=["q_proj", "k_proj", "v_proj", "o_proj", "gate_proj", "up_proj", "down_proj"],
    inference_mode=False,  # 训练模式
    r=8,  # Lora 秩
    lora_alpha=32,  # Lora alaph，具体作用参见 Lora 原理
    lora_dropout=0.1  # Dropout 比例
)
 
base = AutoModelForCausalLM.from_pretrained(model_path, torch_dtype=torch.bfloat16, low_cpu_mem_usage=True)
base_tokenizer = AutoTokenizer.from_pretrained(model_path)
lora_model = PeftModel.from_pretrained(
    base,
    lora_path,
    torch_dtype=torch.float16,
    config=config
)
model = lora_model.merge_and_unload()
model.save_pretrained(output_path)
base_tokenizer.save_pretrained(output_path)
