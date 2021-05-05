let press= document.getElementById("sort");
let div = document.getElementById("filtre");

press.addEventListener('click', () =>{
   if(div.style.display === 'none'){
      div.style.display = 'block';
   }else {
      div.style.display = 'none';
   }
})


const creations = document.querySelectorAll(".creation")

console.log(creations)