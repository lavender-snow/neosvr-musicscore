document.addEventListener('DOMContentLoaded', function(e) {
    let addColumnButton = document.getElementById('add-column-button');
    addColumnButton.addEventListener('click', function(e) {
        addColumn();
    })
    for (let i = 0; i < 30; i++){
        addColumn();
    }
});

function addColumn() {
    let table = document.getElementById("score-table");
    let scales = ['','A0','A#0','B0','C1','C#1','D1','D#1','E1','F1','F#1','G1','G#1','A1','A#1','B1','C2','C#2','D2','D#2','E2','F2','F#2','G2','G#2','A2','A#2','B2','C3','C#3','D3','D#3','E3','F3','F#3','G3','G#3','A3','A#3','B3','C4','C#4','D4','D#4','E4','F4','F#4','G4','G#4','A4','A#4','B4','C5','C#5','D5','D#5','E5','F5','F#5','G5','G#5','A5','A#5','B5','C6','C#6','D6','D#6','E6','F6','F#6','G6','G#6','A6','A#6','B6','C7','C#7','D7','D#7','E7','F7','F#7','G7','G#7','A7','A#7','B7','C8']
    let columnLength = table.rows[0].cells.length+1;
    for (let i = 0; i < table.rows.length; i++) {
        let cell = table.rows[i].insertCell(-1);
        if (i == 0){
            cell.innerHTML = `${columnLength}`;
        }else{
            cell.innerHTML = `<input type="checkbox" value="${scales[i]}_${columnLength}">`;
        }
    }
}