<script>
    export let profiles
    export let activeProfile = null
    function handleClick(event){
        console.log(event.target.id)
        for (let i = 0; i < profiles.length; i++) {
            if(profiles[i].id === event.target.id){
                activeProfile = profiles[i]
            }
        }
    }

</script>

{#each profiles as profile}

    <button id={profile.id} on:click={handleClick}>{profile.name}</button>

{/each}

<style>

</style>