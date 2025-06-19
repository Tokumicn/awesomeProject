from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import mlx.core as mx
from mlx_lm import load, generate

app = FastAPI(title="Qwen3-0.6B-MLX API")

# 加载模型
model_name = "Qwen/Qwen3-0.6B-MLX-4bit"
model, tokenizer = load(model_name)

class TextRequest(BaseModel):
    prompt: str
    max_tokens: int = 100

@app.post("/generate")
async def generate_text(request: TextRequest):
    try:
        response = generate(
            model,
            tokenizer,
            prompt=request.prompt,
            max_tokens=request.max_tokens,
            verbose=True,
        )
        
        return {"generated_text": response}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@app.get("/")
async def root():
    return {"message": "Welcome to Qwen3-0.6B-MLX API. Use /generate endpoint to generate text."}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000) 