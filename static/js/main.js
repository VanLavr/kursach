const buttons = document.querySelectorAll(".first__button");
console.log(buttons)

const negrFuncton = () => {
    console.log(123);
}

for (let button of buttons) {
    button.addEventListener("click", () => negrFuncton())
}