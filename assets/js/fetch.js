document.addEventListener('DOMContentLoaded', function(e) {
    createScoreButton = document.getElementById("create-score-button");
    createScoreButton.addEventListener('click',function(e){
        createScore();
    });
})

function createScore() {
    let form = document.getElementById("score-form");
    let formData = new FormData(form);
    fetch("/score-creator/create", {
        method: "post",
        headers: {},
        body: formData
    }).then(function (response){
        resultDiv = document.getElementById("fetch-result");
        if (response.status === 200) {
            response.text().then(hash => {
                console.log(hash);
                let url = `${window.location.protocol}//${window.location.host}/score-creator/score?v=${hash}`;
                resultDiv.innerHTML = `<p>URLの生成が完了しました。GenerateScoreMachineに以下のURLを入力してください。<br>${url}</p>`;
            });
        }else if(response.status === 503){
            resultDiv.innerHTML = `<p class='text-danger'>${response.body.json()}</p>`;
        }else{
            resultDiv.innerHTML = "<p class='text-danger'>原因不明のエラーが発生しました</p>";
        }
    }).catch(function (response){
        console.log(response);
    });
}