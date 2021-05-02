export function display() {
    fetch('/api')
        .then(function (response) {
            return response.json();
        })
        .then(function (data) {
            sessionStorage.setItem('datasave', data)
            document.getElementById("nuke").onclick = function () {
                document.location.reload()
                sessionStorage.clear()
            }
            document.getElementById("srch").onclick = function () {
                srch(data)
            };
            appendData(data);

        })

    function appendData(heroes) {
        let table = document.getElementById("tab");
        let row = table.insertRow();
        var row1 = table.insertRow();
        // Mise en place du nombre de page par défaut qui est de 20
        for (let i = 0; i < heroes.length; i++) { // on va parcourir le tableau d'hero
            // on creer une variable qui va prendre pour information tout ce qui ce trouve dans l'id "tab" // on creer un tableau
            let cell = row.insertCell();

            console.log("tourne");
            if (i === 10) {
                cell = row1.insertCell();
                let icon = row.insertCell();
                cell.innerHTML = heroes[i].name;
                let imageName = ["background-image:url('", heroes[i].image, "')"]; // on ajoute les icones des personnages dans le style de la page
                imageName = imageName.join("");
                icon.style = imageName;

            }

            cell.innerHTML = heroes[i].name;
        }

    }
}
export function srch(heroes) {
    let research = document.getElementById("searchInput").value;
    let search = new RegExp(research);
    let finded = []
    let valid = []
    console.log(research)
    for (let i = 0; i < heroes.length; i++) {
        valid.push(false)
        for (const property in heroes[i]) {
            if (typeof heroes[i][property] == 'object') {
                for (const subproperty in heroes[i][property]) {
                    if (typeof heroes[i][property][subproperty] == 'object') {
                        for (const sublist in heroes[i][property][subproperty]) {
                            if (search.test(heroes[i][property][subproperty][sublist]) == true) {
                                valid[i] = true
                            }
                        }
                    } else {
                        if (search.test(heroes[i][property][subproperty]) == true) {
                            valid[i] = true
                        }
                    }
                }
            } else {
                if (search.test(heroes[i][property]) == true) {
                    valid[i] = true
                }
            }
        }
        if (valid[i] == true) {
            finded.push(heroes[i])
        }
    }
    finded = JSON.stringify(finded)
    sessionStorage.setItem("data", finded)
}