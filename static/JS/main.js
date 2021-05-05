const charactersList = document.getElementById('charactersList');
const searchinput = document.getElementById('searchInput');
let persons = [];

searchinput.addEventListener('keyup', (e) => {
    const searchString = e.target.value.toLowerCase();

    const filteredCharacters = persons.filter((character) => {
        return (
            character.name.toLowerCase().includes(searchString) ||
            character.house.toLowerCase().includes(searchString)
        );
    });
    displayCharacters(filteredCharacters);
});

const loadCharacters = async () => {
    try {
        const res = await fetch('https://groupietrackers.herokuapp.com/api/artists');
        persons = await res.json();
        displayCharacters(persons);
    } catch (err) {
        console.error(err);
    }
};

const displayCharacters = (characters) => {
    const htmlString = characters
        .map((character) => {
            return `
            <li class="character">
                <h2>${character.name}</h2>
                <p>House: ${character.house}</p>
                <img src="${character.image}"></img>
            </li>
        `;
        })
        .join('');
    charactersList.innerHTML = htmlString;
};

searchinput.addEventListener('keyup', function(){
    const input = searchinput.value;

    const result = persons.filter(item => item.name.toLowerCase().includes(input.toLowerCase()));

    let suggestion = '';

    if(input !=''){
    result.forEach(resultItem => 
            suggestion +=`
                <div class="suggestion">${resultItem.name}</div>
            `
        )
    }
    document.getElementById("suggestions").innerHTML = suggestion;

})

console.log(persons)
loadCharacters();