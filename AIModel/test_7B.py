import os
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM
from safetensors.torch import load_file as load_safetensors

pretrained_model_path = ""  # 预训练模型所在目录
finetuned_weights_dir = ""  # 微调后权重所在目录

merged_file = os.path.join(finetuned_weights_dir, "model.safetensors")

# 加载预训练模型结构和 tokenizer
print("加载预训练模型结构和 tokenizer...")
tokenizer = AutoTokenizer.from_pretrained(pretrained_model_path, trust_remote_code=True)
model = AutoModelForCausalLM.from_pretrained(pretrained_model_path, trust_remote_code=True)

# 加载微调权重
if os.path.exists(merged_file):
    print("加载微调权重文件：", merged_file)
    ft_state_dict = load_safetensors(merged_file)  # 更新模型参数，strict=False 允许部分键不匹配（建议观察警告信息）
    model.load_state_dict(ft_state_dict, strict=False)
else:
    print("未找到合并后的权重文件 model.safetensors，请确认是否已合并权重。")
    exit(1)

# 将模型传入 GPU 并支持多卡
if torch.cuda.is_available():
    device = torch.device("cuda")  # 默认使用 GPU
else:
    device = torch.device("cpu")  # 如果没有 GPU，使用 CPU

# 使用 DataParallel 支持多卡（如果超过1张卡可用）
if torch.cuda.device_count() > 1:
    print(f"使用 {torch.cuda.device_count()} 个 GPU")
    model = torch.nn.DataParallel(model)  # 包装模型，使其跨多个 GPU 并行

model.to(device)
model.eval()

prompt = ''
while True:
    text = input("请输入问题或信息（输入 '-1' 退出）：")
    if text == '-1':
        break

    inputs = tokenizer(prompt + text, return_tensors="pt").to(device)

    with torch.no_grad():
        outputs = model.module.generate(**inputs, max_new_tokens=1000000)  # 限制生成长度以防止显存不足

    # 只解码生成的部分，排除用户的输入和 prompt
    generated_part = outputs[0][len(inputs.input_ids[0]):]
    result = tokenizer.decode(generated_part, skip_special_tokens=True)
    
    print("生成结果：", result)