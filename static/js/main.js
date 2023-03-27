async function getJSON() {
    const response = await fetch("http://localhost:8080");
    return response.json();
}

console.log(getJSON());