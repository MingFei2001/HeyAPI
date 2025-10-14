// return alert and message upon button click
document.addEventListener("DOMContentLoaded", () => {
  const button = document.getElementById("myButton");
  const message = document.getElementById("message");

  button.addEventListener("click", () => {
    message.textContent = "Button clicked! JavaScript is working.";
    alert("Hello from the other side!");
  });
});
