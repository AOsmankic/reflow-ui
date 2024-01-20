<script>
    import Chart from 'chart.js/auto';
    import {onMount} from "svelte";
    export let activeProfile
    let myChart
    let portfolio
    let ctx
    let activeName = ""
    let tempCurve = [];
    export function updatePreview(selectedProfile) {
        if (selectedProfile == null){
            return
        }
        console.log(selectedProfile)
        tempCurve = []
        activeName = selectedProfile.name

        for (let i = 0; i < selectedProfile.values.length; i++) {
            let selectedValue = selectedProfile.values[i]
            tempCurve.push({x: selectedValue.X, y: selectedValue.Y})
        }
        myChart.data.datasets[0].data = tempCurve
        myChart.update()

    }

    onMount(() => {
        ctx = portfolio.getContext('2d')
        myChart = new Chart(ctx, {
            data:{
                datasets: [{
                    type: 'line',
                    label: 'Expected Temperature',
                    data: tempCurve,
                }]
            },
            options: {
                animation: false,
                scales: {
                    y: {
                        beginAtZero: true
                    },
                    x: {
                        beginAtZero: true,
                        type: 'linear',
                    },
                },

            }
        })
    })
</script>
<div class="preview-container">
    <h1 class="profile-name">{activeName}</h1>
    <canvas bind:this={portfolio} width={400} height={400} />
</div>


<style>
    .profile-name {
        font-size: 56px;
        color: #1b2636;
    }
    .preview-container {
        width: 50%;
    }
</style>