from flask import Flask, request, jsonify
import time

app = Flask(__name__)

@app.route('/predict', methods=['POST'])
def predict():
    data = request.json
    text = data.get('text')
    if not text:
        return jsonify({"error": "No text provided"}), 400

    # 模拟推理过程
    time.sleep(1)
    return jsonify({"text": text, "intent": "example_intent"})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)