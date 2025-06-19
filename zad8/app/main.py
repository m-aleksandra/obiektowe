from flask import Flask, render_template, request, jsonify

app = Flask(__name__)

@app.route("/")
def index():
    return render_template("index.html")

@app.route("/hello", methods=["POST"])
def hello():
    data = request.get_json()
    name = data.get("name")
    if not name:
        name = "stranger"
    return jsonify({"message": f"Hello, {name}!"})

@app.route("/user/<name>")
def user(name):
    return f"User page for: {name}"

if __name__ == "__main__":
    app.run(debug=True)
