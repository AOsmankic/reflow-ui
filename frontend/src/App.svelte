<script>
    import logo from './assets/images/logo-universal.png'
    import Chart from 'chart.js/auto';
    import Bottombar from "./Bottombar.svelte";
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
    let bottomOpen = false
    let updatedProfile


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
            open=false
        }else if(event.detail.direction === "right"){
            open=true
        }else if(event.detail.direction === "bottom"){
            bottomOpen=false
        }else if(event.detail.direction === "top"){
            bottomOpen=true
        }

        console.log(event)
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

    function handleKeyPress(event) {
        console.log(event)
    }

    let clear
    $: {
        clearInterval(clear)
        clear = setInterval(getData, 1000)
    }
    $: updatedProfile, getData()

    function closeMenus() {
        open=false
    }

</script>
<Sidebar bind:open bind:updatedProfile/>
<Bottombar bind:bottomOpen/>
<Navbar bind:sidebar={open}/>
<div class="graph" use:swipe={{ timeframe: 300, minSwipeDistance: 100, touchAction: 'pan-y' }} on:swipe={graphSwipeHandler}>
    <canvas bind:this={portfolio} />
</div>
<svelte:head>
    <link href="https://unpkg.com/tailwindcss@^1.0/dist/tailwind.min.css" rel="stylesheet"/>
</svelte:head>

<style>

    :global(body) {
        padding: 0;

        overflow-y: hidden;
    }

    .graph{
        height: 100%;
        padding-bottom: 10%;
    }

    :global(.splitpanes__pane) {
        box-shadow: 0 0 3px rgba(0, 0, 0, .2) inset;
        justify-content: center;
        align-items: center;
        display: flex;
        position: relative;
        overflow-y: auto;
    }



</style>
