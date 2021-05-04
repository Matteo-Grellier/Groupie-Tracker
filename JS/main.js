const loadData = artists => {
    console.log(artists)
}

fetch("https://cors-anywhere.herokuapp.com/https://groupietrackers.herokuapp.com/api/artists")
.then((response) => response.json())
.then(loadData)


const searchinput = document.getElementById('searchInput');

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