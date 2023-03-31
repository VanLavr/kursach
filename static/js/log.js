const fetchProducts = async () => {
    const response = await fetch("http://localhost:8080/all");

    if (response.ok) {
        return (json = await response.json());
    } else {
        console.log("Ошибка HTTP: " + response.status);
    }
};

const buttons = document.querySelectorAll(".button");

buttons[0].addEventListener("click", () => console.log(fetchProducts()));
