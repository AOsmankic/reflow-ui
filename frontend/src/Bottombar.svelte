<script>
    import {swipe} from "svelte-gestures";
    import {
        GetState
    } from '../wailsjs/go/main/App.js'
    export let bottomOpen = false
    let clear
    let currentState = {
        Temp: -1,
        SelectedProfile: "NaN",
        SecondsLeft: -1
    }

    function handleDrag(event) {
        if(event.detail.direction === "bottom"){
            bottomOpen=false
        }else if(event.detail.direction === "top"){
            bottomOpen=true
        }
    }
    function fetchState() {
        GetState().then((s) => {
            console.log(s)
            currentState = s
        })
    }
    function prettifyTime(seconds){
        let hours = 0
        if (seconds < 1){
            return "N/A"
        }
        while (seconds >= 60){
            seconds -= 60
            hours++
        }
        let returnStr = ""
        if (hours > 0){
            returnStr += hours + " Hours & "
        }
        return returnStr + seconds + " seconds"
    }

    $: {
        clearInterval(clear)
        clear = setInterval(fetchState, 1000)
    }

</script>

<aside class="absolute w-full h-full bg-black-200 border-r-2 shadow-lg" class:bottomOpen use:swipe={{ timeframe: 300, minSwipeDistance: 100, touchAction: 'pan-y' }} on:swipe={handleDrag}>
    <div class="div-container">
        <svg width=32 height=24 on:click={() => {bottomOpen = !bottomOpen}}>
            <line id="top" x1=0 y1=12  x2=32 y2=12 />
            <line id="middle" x1=0 y1=22 x2=32 y2=22 />
            <line id="bottom" x1=0 y1=32 x2=32 y2=32 />
        </svg>
        <div class="grid-container">
            <div class="grid-item status-cell">Current Temperature<br/>{currentState.Temp}</div>

            <div class="grid-item status-cell">Selected Profile<br/>{currentState.SelectedProfile}<br/></div>
            <div class="grid-item status-cell">Estimated Time Left<br/>{prettifyTime(currentState.SecondsLeft)}</div>
        </div>
    </div>
</aside>
<style>
    aside {
        bottom: -95%;
        transition: bottom 0.3s ease-in-out;
    }

    .div-container{
        background-color: #f2f2f2;
        border-radius: 30px;
        border: 2px solid #0e1b25;
        height: 100%;
        color: #0e1b25;
        display: block;
        align-content: center;
        justify-content: center;
    }

    .headers{
        font-size: 56px;
    }

    .bottomOpen {
        bottom: -50%
    }

    .status-cell {
        height: 60px;
        display: flex;
        justify-content: center;
        align-content: center;
        border-radius: 20px;
        background-color: #3a4d72;
        color: #f2f2f2;
        padding: 14px 28px;
        height: 95%;
        width: 95%;
        margin: 10px;
    }
    .grid-container {
        display: grid;
        grid-template-columns: auto auto auto;
        padding: 10px;
        width: 100%;
        margin-bottom: 100px;
        margin-top: 30px;
    }
    .grid-item {
        border: 1px solid rgba(0, 0, 0, 0.8);
        padding: 20px;
        font-size: 30px;
        text-align: center;
    }
    svg {
        min-height: 24px;
        margin-left: 50%;

    }

    svg line {
        stroke: currentColor;
        stroke-width: 3;
        align-content: center;
    }
    .top {
        transform: translate(6px, 0px) rotate(45deg)
    }

    .middle {
        opacity: 0;
    }

    .bottom {
        transform: translate(-12px, 9px) rotate(-45deg)
    }



</style>