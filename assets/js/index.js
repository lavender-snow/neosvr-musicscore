const scales = ['C8','B7','A#7','A7','G#7','G7','F#7','F7','E7','D#7','D7','C#7','C7','B6','A#6','A6','G#6','G6','F#6','F6','E6','D#6','D6','C#6','C6','B5','A#5','A5','G#5','G5','F#5','F5','E5','D#5','D5','C#5','C5','B4','A#4','A4','G#4','G4','F#4','F4','E4','D#4','D4','C#4','C4','B3','A#3','A3','G#3','G3','F#3','F3','E3','D#3','D3','C#3','C3','B2','A#2','A2','G#2','G2','F#2','F2','E2','D#2','D2','C#2','C2','B1','A#1','A1','G#1','G1','F#1','F1','E1','D#1','D1','C#1','C1','B0','A#0','A0']
const scales_en = ["C","D","E","F","G","A","B"]
const scales_jp = ["ド","レ","ミ","ファ","ソ","ラ","シ"]

document.addEventListener('DOMContentLoaded', function(e) {
    addColumn(1);

    let maxScaleRange = document.getElementById("max-scale-range");

    for (let i = 0; i < scales.length-1; i++){
        let option = document.createElement("option");
        option.value = scales[i];
        option.text = scales[i];
        maxScaleRange.appendChild(option);
    }

    let minScaleRange = document.getElementById("min-scale-range");

    minScaleRange.addEventListener('change', function(e){
        changeTable();
    })

    maxScaleRange.selectedIndex = scales.indexOf("A5");
    changeMinScale(scales.indexOf("A5"));
    minScaleRange.selectedIndex = 20;

    maxScaleRange.addEventListener('change', function(e) {
        let maxScaleSelectedIndex = maxScaleRange.selectedIndex;
        let minScaleRange = document.getElementById("min-scale-range");
        if (maxScaleSelectedIndex >= scales.indexOf(minScaleRange.options[minScaleRange.selectedIndex].value)){
            changeMinScale(maxScaleSelectedIndex)
        }

        changeTable();
    });

    let scaleLabelsSwitch = document.getElementById("scale-labels-switch");
    scaleLabelsSwitch.addEventListener('change', function(e) {
        if (scaleLabelsSwitch.checked) {
            changeScalesLabel(scales_en, scales_jp);
        } else {
            changeScalesLabel(scales_jp, scales_en);
        }
    });

    changeTable();
});

function changeMinScale(maxScaleSelectedIndex){
    let minScaleRange = document.getElementById("min-scale-range");

    while (minScaleRange.firstChild) {
        minScaleRange.removeChild(minScaleRange.firstChild);
    }

    for (let i = maxScaleSelectedIndex+1; i < scales.length; i++) {
        let option = document.createElement("option");
        option.value = scales[i];
        option.text = scales[i];
        minScaleRange.appendChild(option);
    }
    minScaleRange.selectedIndex = 0;
}

function addColumn(index) {
    let table = document.getElementById("score-table");
    let columnLength = table.rows[0].cells.length+1;

    if (columnLength === index){
        for (let i = 0; i < table.rows.length; i++) {
            let cell = table.rows[i].insertCell(-1);
            if (i === 0){
                cell.innerHTML = `${columnLength}`;
            }else{
                let input = document.createElement("input");
                let scoreId = `${scales[i-1]}_${columnLength}`;
                input.classList.add("form-check-input");
                input.type = "checkbox";
                input.name = "score";
                input.value = scoreId;
                input.addEventListener('click', function(e){
                    addColumn(index+1);
                })
                cell.appendChild(input);
            }
        }
        table.parentElement.parentElement.scrollLeft = table.parentElement.parentElement.scrollWidth;
    }
}

function changeTable(){
    let maxScaleRange = document.getElementById("max-scale-range");
    let minScaleRange = document.getElementById("min-scale-range");
    let scaleTable = document.getElementById("scale-table");
    let scoreTable = document.getElementById("score-table");

    maxScaleIndex = maxScaleRange.selectedIndex;
    minScaleIndex = scales.indexOf(minScaleRange.options[minScaleRange.selectedIndex].value)+1;

    for (let i = 1; i < scaleTable.rows.length; i++){
        if (maxScaleIndex < i && i <= minScaleIndex) {
            scaleTable.rows[i].style.display = "table-row";
            scoreTable.rows[i].style.display = "table-row";
        } else {
            scaleTable.rows[i].style.display = "none";
            scoreTable.rows[i].style.display = "none";
        }

    }
}

function changeScalesLabel(beforeArr, afterArr){
    let scaleTable = document.getElementById("scale-table");
    for (let i = 0; i < scaleTable.rows.length; i++){
        scaleTable.rows[i].innerHTML = beforeArr.reduce((a, c, i) => {
            const reg = new RegExp(beforeArr[i], 'gm')
            c = a.replace(reg, afterArr[i])
            return c
        }, scaleTable.rows[i].innerHTML)
    }
}