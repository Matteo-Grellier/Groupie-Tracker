
const searchinput = document.getElementById('searchInput');
let groupesCharacters = [];

searchinput.addEventListener('keyup', (e) => {
    const searchString = e.target.value.toLowerCase();
    const filteredGroupes = groupesCharacters.filter((artist) => {
        return (
            artist.name.includes(searchString)
        );
    });
    displayCharacters(filteredGroupes);
});

const loadCharacters = async () => {
    try {
        const res = await fetch('/api');
        groupesCharacters = await res.json();
        displayCharacters(groupesCharacters);
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
                <img src="${character.image}"></img>
            </li>
        `;
        })
        .join('');
    charactersList.innerHTML = htmlString;
};

searchinput.addEventListener('keyup', function(){
    
    const input = searchinput.value.toLowerCase();

    const result = groupesCharacters.filter(item => item.name.toLowerCase().includes(input.toLowerCase()));

    let suggestion = '';

    if(input !=''){
    result.forEach(resultItem => 
            suggestion +=`<div class="suggestion">${resultItem.name}</div>`
        )
    }
    document.getElementById("suggestions").innerHTML = suggestion;

})

loadCharacters();