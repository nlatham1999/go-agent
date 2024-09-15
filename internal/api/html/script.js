
//function to call the /setup endpoint
function setup() {

    console.log("made setup request");

    fetch('/setup', {
        method: 'POST',
    })
        .then(response => response.json())
        .then((data) => {
            console.log("made setup request");
        });
}

// function to call the go endpoint (post) and print the status
function go() {
    fetch('/go', {
        method: 'POST'
    })
        .then(response => response.json())
        .then((data) => {
            console.log(data.status);
        });
}

//funcion to call the model endpoint and get the data
function model() {

    console.log("made model request");

    fetch('/model')
        .then(response => response.json())
        .then((data) => {
            processData(data);
        });
}

function processData(data) {

    console.log(document.getElementById('min-pxcor'));
    console.log(document.getElementById('min-pxcor"').innerText)
    console.log(data);


    //loop through the data[patches] array
    minX = data.patches[0].x;
    minY = data.patches[0].y;
    maxX = data.patches[0].x;
    maxY = data.patches[0].y;
    data.patches.forEach((patch) => {
        // get the x and y coordinates of the patch
        let x = patch.x;
        let y = patch.y;

        // get the min and max x and y values
        if (x < minX) {
            minX = x;
        }
        if (x > maxX) {
            maxX = x;
        }
        if (y < minY) {
            minY = y;
        }
        if (y > maxY) {
            maxY = y;
        }
    });

    // set the divs to the max and min values
    document.getElementById('min-pxcor"').innerText = minX;
    document.getElementById('min-pycor"').innerText = minY;
    document.getElementById('max-pxcor"').innerText = maxX;
    document.getElementById('max-pycor"').innerText = maxY;
}