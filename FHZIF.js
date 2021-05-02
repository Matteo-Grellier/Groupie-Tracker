const concerts = [
    {name: 'France', ville:1},
    {name: ' USA', ville:2},
    {name: 'Canada', ville:3}
  ];
  
  const searchinput = document.getElementById("searchInput");
  
  searchinput.addEventListener('keyup', function(){
      
    const input = searchinput.value.toLowerCase();
  
    const result = concerts.filter(item => item.name.toLowerCase().includes(input.toLowerCase()));
  
    let suggestion = '';
  
    if(input !=''){
    result.forEach(resultItem => 
            suggestion +=`<div class="suggestion">${resultItem.name}</div>`
        )
    }
    document.getElementById("suggestions").innerHTML = suggestion;
  
  })


/*
  <div id="search_location">
  <input class="input" id="searchInput" type="text" placeholder="Search concert..." />
  <button class="btn" type="submit">Search</button>
  <div id="suggestions"></div>
</div> */