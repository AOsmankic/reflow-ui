<script>
    import logo from './assets/images/logo-universal.png'
    import Chart from 'chart.js/auto';

    import {
        GetTempHistory,
        FindMaxY,
        FindMaxX,
        GetAllProfiles,
        Init,
        GetSelectedProfile,
        ShiftTempBuffer,
        ResetBuffer,
    } from '../wailsjs/go/main/App.js'
    import {LogPrint} from "../wailsjs/runtime/runtime.js";
    import {onMount} from "svelte";
    import {swipe} from "svelte-gestures";
    import Navbar from './Navbar.svelte'
    import Sidebar from './Sidebar.svelte'

    let profiles
    let portfolio;
    let points = []
    let tempBuffer = []
    let tempCurve = []

    let selectedProfile
    let myChart
    let ctx
    let open = false

    onMount(async (promise) => {
        let tmpTempCurve
        Init().then(() => GetAllProfiles().then(p => profiles = p)).then(() => console.log(profiles)).then(() => GetSelectedProfile().then(p => selectedProfile = p).then(() => tmpTempCurve = selectedProfile.values)).then(() => {
            tempCurve = [];
            for (let tempCurveKey in tmpTempCurve) {
                let tempCoord = tmpTempCurve[tempCurveKey]
                tempCurve.push(
                    {x: tempCoord.X, y: tempCoord.Y}
                )
            }

            myChart.data.datasets[0].data = points
            myChart.data.datasets[1].data = tempCurve
            myChart.update()
        })
        ctx = portfolio.getContext('2d');
        // Initialize chart using default config set
        myChart = new Chart(ctx, {
            data:{
                datasets: [{
                    type: 'scatter',
                    label: 'Real Temperature',
                    data: points
                }, {
                    type: 'line',
                    label: 'Expected Temperature',
                    data: tempCurve,
                }],
            },
            options: {
                animation: false,
                scales: {
                    y: {
                        beginAtZero: true
                    },
                    x: {
                        beginAtZero: true,
                        min: 0
                    }
                }
            }
        })
    })

    function graphSwipeHandler(event) {
        if (event.detail.direction === "left"){
            shiftTempLeft()
        }else if(event.detail.direction === "right"){
            shiftTempRight()
        }
    }

    function shiftTempLeft(){
        ShiftTempBuffer(-10).then(() =>{
            getData()
        })
    }
    function shiftTempRight(){
        ShiftTempBuffer(10).then(() =>{
            getData()
        })
    }

    async function resetTempCapture(){
        await ResetBuffer()
        getData()
    }

    async function getData() {
        GetTempHistory().then(buffer => tempBuffer = buffer).then(() => {

            points = []
            for (let tempBufferKey in tempBuffer) {
                let bufferCoord = tempBuffer[tempBufferKey]
                points.push(
                    {x: bufferCoord.X, y: bufferCoord.Y}
                )
            }
            myChart.data.datasets[0].data = points

        }).then(() => {myChart.update()})
        let tmpTempCurve
        GetSelectedProfile().then(p => selectedProfile = p).then(() => tmpTempCurve = selectedProfile.values).then(() => {
            tempCurve = [];
            for (let tempCurveKey in tmpTempCurve) {
                let tempCoord = tmpTempCurve[tempCurveKey]
                tempCurve.push(
                    {x: tempCoord.X, y: tempCoord.Y}
                )
            }
            myChart.data.datasets[1].data = tempCurve
        }).then(myChart.update())

        // myChart.update()
    }

    let clear
    $: {
        clearInterval(clear)
        clear = setInterval(getData, 10000)
    }
</script>
<Sidebar bind:open/>
<Navbar bind:sidebar={open}/>
<div use:swipe={{ timeframe: 300, minSwipeDistance: 100, touchAction: 'pan-y' }} on:swipe={graphSwipeHandler} >
    <canvas bind:this={portfolio} width={1920} height={1080} />
<!--    <button on:click={shiftTempLeft}>Shift Left</button>-->
<!--    <button on:click={shiftTempRight}>Shift Right</button>-->
<!--    <button on:click={resetTempCapture}>Reset</button>-->
</div>
<svelte:head>
    <link href="https://unpkg.com/tailwindcss@^1.0/dist/tailwind.min.css" rel="stylesheet"/>
</svelte:head>

<style>

    :global(body) {
        padding: 0;
    }
    #logo {
        display: block;
        width: 50%;
        height: 50%;
        margin: auto;
        padding: 10% 0 0;
        background-position: center;
        background-repeat: no-repeat;
        background-size: 100% 100%;
        background-origin: content-box;
    }

    .result {
        height: 20px;
        line-height: 20px;
        margin: 1.5rem auto;
    }

    .input-box .btn {
        width: 60px;
        height: 30px;
        line-height: 30px;
        border-radius: 3px;
        border: none;
        margin: 0 0 0 20px;
        padding: 0 8px;
        cursor: pointer;
    }

    .input-box .btn:hover {
        background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
        color: #333333;
    }

    .input-box .input {
        border: none;
        border-radius: 3px;
        outline: none;
        height: 30px;
        line-height: 30px;
        padding: 0 10px;
        background-color: rgba(240, 240, 240, 1);
        -webkit-font-smoothing: antialiased;
    }

    .input-box .input:hover {
        border: none;
        background-color: rgba(255, 255, 255, 1);
    }

    .input-box .input:focus {
        border: none;
        background-color: rgba(255, 255, 255, 1);
    }

    :global(.splitpanes__pane) {
        box-shadow: 0 0 3px rgba(0, 0, 0, .2) inset;
        justify-content: center;
        align-items: center;
        display: flex;
        position: relative;
    }

</style>
