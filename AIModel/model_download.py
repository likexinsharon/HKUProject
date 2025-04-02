from modelscope import snapshot_download 
save_path = ""
model_dir = snapshot_download('deepseek-ai/deepseek-llm-7b-chat', cache_dir=save_path, revision='master')
