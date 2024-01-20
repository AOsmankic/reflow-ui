<script>
    import {onMount} from "svelte";

    export let profiles
    export let activeProfile = null
    let selectedButton
    function handleClick(event){
        console.log(event.target.id)
        for (let i = 0; i < profiles.length; i++) {
            if(profiles[i].id === event.target.id){
                activeProfile = profiles[i]
            }
        }
        selectedButton = event.target.id
    }
    function updateSelection(){
        if (activeProfile == null){
            return
        }
        selectedButton= activeProfile.id
    }

    $: activeProfile, updateSelection()


</script>

{#each profiles as profile}

    <button class="profile-button" class:active={selectedButton === profile.id} id={profile.id} on:click={handleClick}>{profile.name}</button>

{/each}

<style>
    .profile-button {
        border: 2px solid #1b2636;
        border-radius: 9px;
        background-color: #3a4d72;
        padding: 14px 28px;
        font-size: 16px;
        width: 90%;
        margin-top: 15px;
        text-align: right;
    }
    .active {
        background-color: #0e1b25;
    }
</style>