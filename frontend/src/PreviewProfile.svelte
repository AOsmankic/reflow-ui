<script>
    import Chart from 'chart.js/auto';
    import {onMount} from "svelte";
    export let activeProfile
    let myChart
    let portfolio
    let ctx
    export function updatePreview() {
        let tempCurve = [];
        for (let i = 0; i < activeProfile.values.length; i++) {
            let selectedValue = activeProfile.values[i]
            console.log(selectedValue)
            tempCurve.push({x: selectedValue.X, y: selectedValue.Y})
        }
        console.log(tempCurve)
        myChart = new Chart(ctx, {
            data:{
                datasets: [{
                    type: 'line',
                    label: 'Expected Temperature',
                    data: tempCurve,
                },
                {
                    type: 'scatter',
                    label: 'Real Temperature',
                    data: []
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
                        max: 40
                    }
                }
            }
        })
    }

    onMount(() => {
        ctx = portfolio.getContext('2d')
        // updatePreview()
    })
</script>

<canvas bind:this={portfolio} width={400} height={400} />

