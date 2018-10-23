function greeter(name) {
  return "Hello! " + name
}

const user = "Sreekar"

document.getElementById("MessageBox").innerText = greeter(user)