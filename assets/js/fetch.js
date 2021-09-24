document.addEventListener('DOMContentLoaded', function(e) {
    createScoreButton = document.getElementById("create-score-button");
    createScoreButton.addEventListener('click',function(e){
        createScore();
    });
})

function createScore() {
    let form = document.getElementById("score-form");
    let formData = new FormData(form);
    for (let value of formData) {
        console.log(value);
    }
    fetch("/score-creator/create", {
        method: "post",
        headers: {},
        body: formData
    }).then(function (response){
        if (response.status === 200) {
            console.log("Suceess");
            console.log(response.body);
        }else{
            console.log("Failed");
        }
    }).catch(function (response){
        console.log(response);
    });
}