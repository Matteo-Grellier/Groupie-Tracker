/* Luke */
// exemple :

// recherche interactive
const persons = [ // il faudra mettre les artist etc ...

    {name:'Luke', age: 18 },
    {name:'Loic', age: 21 },
    {name:'Yann', age: 19 },
    {name:'Nicolas', age: 25 },
    {name:'Matteo', age: 32 },
    {name:'Paul', age: 20 },
    {name:'Jeff', age: 12 }
];

const searchinput = document.getElementById('searchInput');
 
searchinput.addEventListener('keyup', function(){
    const input = searchinput.value;

    const result = persons.filter(item => item.name.toLowerCase().includes(input.toLowerCase()));

    let suggestion = '';

    if(input != '' ){
    result.forEach(resultItem => 
            suggestion +=`
                <div class="suggestion">${resultItem.name}</div>
            `
        )
    }
    document.getElementById("suggestions").innerHTML = suggestion;

})

/* Yann */

/* Nicolas */

/* Matteo */