// fetch the random number API as data
document.addEventListener("DOMContentLoaded", () => {
  fetch("/api/random")
    .then((res) => res.json())
    .then((data) => {
      document.getElementById("randomMsg").textContent =
        "Random number: " + data.random;
    })
    .catch((err) => {
      document.getElementById("randomMsg").textContent =
        "Error fetching random number";
      console.error(err);
    });
});
