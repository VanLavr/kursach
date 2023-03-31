const buttons = document.querySelectorAll(".first__button");
console.log(buttons)

const negrFuncton = () => {
    console.log(123);
}

for (let button of buttons) {
    button.addEventListener("click", () => negrFuncton())
}
async function getJSON() {
    const response = await fetch("http://localhost:8080");
    return response.json();
}

console.log(getJSON());
