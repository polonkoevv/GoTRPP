function addActors(data){
    let key;
    var target = document.querySelector('.actors');
    target.innerHTML += `
        <div class="actor">
            <img src="${data.posterUrl}" alt="Тут должно быть фото" class="img">
            <form action="/actors/${data.personId}" method="POST">
                <button style="text-decoration: underline;" type="submit" class="description">${data.nameRu}</button>
            </form>
        </div>
    `
}

            // <p><span class="bold">Лучшие фильмы:</span> <br> ${data.films[0].nameRu} <br> ${data[key].films[1].nameRu} <br> ${data[key].films[2].nameRu} <br>} </p>
function readTextFile(file, callback) {
    var rawFile = new XMLHttpRequest();
    rawFile.overrideMimeType("application/json");
    rawFile.open("GET", file, true);
    rawFile.onreadystatechange = function() {
        if (rawFile.readyState === 4 && rawFile.status == "200") {
            callback(rawFile.responseText);
        }
    }
    rawFile.send(null);
}



function changeActors(start){
    readTextFile("sources/persons.json", function(text){
        var data = JSON.parse(text);
        // console.log(text);
        var list = Object.keys(data)
        // console.log(list.length);
        var len = 10;
        var end = start + len
        if(end > list.length){
            end = start + list.length % len;
        }
        document.querySelector('.actors').innerHTML = "";    
        for (let i = start; i < end; i++) {
            addActors(data[list[i]])        
        }
    });
}

changeActors(0)
