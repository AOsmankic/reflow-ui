<script>
    import { Pane, Splitpanes } from 'svelte-splitpanes';
    import Profile from "./Profile.svelte";
    import PreviewProfile from "./PreviewProfile.svelte";
    import {swipe} from "svelte-gestures";
    import {
        GetAllProfiles,
        SetProfile
    } from '../wailsjs/go/main/App.js'
    import {onMount} from "svelte";
    export let open = false
    export let updatedProfile
    let previewPane
    let oldProfile
    let activeProfile
    let profiles = [{
        "name": "test"
    }]
    async function initializeStuffs(){
        profiles = await GetAllProfiles()
        if (profiles.length > 0){
            activeProfile = profiles[0]
        }
    }

    onMount(async () => {
        await initializeStuffs()
        console.log(activeProfile)
        previewPane.updatePreview(activeProfile)
        let elements = document.getElementsByClassName("splitpanes__pane")
        for (let i = 0; i < elements.length; i++) {
            let element = elements[i]
            if (element.className.includes("scrollable")) {
                element.style["overflow"] = "auto"
            }
        }
        console.log(elements)
    })
    function updatePreview(){
        console.log("updated ")
        console.log(activeProfile)
        console.log("updated ")
        if (previewPane != null) {
            previewPane.updatePreview(activeProfile)
        }
    }

    function handleSubmit(event){
        console.log(event)
        SetProfile(activeProfile.id)
        updatedProfile=activeProfile
        open=false

    }
    function handleDrag(event) {
        if (event.detail.direction === "left"){
            open=false
        }else if(event.detail.direction === "right"){
            open=true
        }
    }

    $:  activeProfile, updatePreview()


</script>
<aside class="absolute w-full h-full bg-black-200 border-r-2 shadow-lg " class:open use:swipe={{ timeframe: 300, minSwipeDistance: 100, touchAction: 'pan-y' }} on:swipe={handleDrag} style="overflow-y: auto">
    <Splitpanes class="default-theme scrollable" style="height: 100%; overflow-y: auto">
        <Pane minSize={30} size={30} class="scrollable" >
            <div class="btn-group scrollable">
                <h1 class="profile-name">Profiles</h1>
                <Profile {profiles} bind:activeProfile={activeProfile}/>
            </div>
        </Pane>
        <Pane minSize="{30}">
            <PreviewProfile {activeProfile} bind:this={previewPane}/>
            <input type="button" value="Select" class="right-button" on:click={handleSubmit}>
        </Pane>
    </Splitpanes>
</aside>
<style>
    aside {
        left: -100%;
        transition: left 0.3s ease-in-out;
    }

    .scrollable {
        overflow: auto;
        height: 80%;
    }

    .open {
        left: -10%
    }
    .profile-name {
        font-size: 56px;
        color: #1b2636;
        text-align: right;
        margin-right: 20px;
    }

    .btn-group button {
        display: block;
        padding: 14px 28px;
        margin-top: 15px;
        overflow-y: auto;
        height: 60%;
    }
    .btn-group {
        width: 100%;
        overflow-y: auto;
    }

    .right-button {
        border: 2px solid #1b2636;
        border-radius: 9px;
        background-color: #3a4d72;
        padding: 14px 28px;
        font-size: 16px;
        width: 200px;
        margin-top: 15px;
        float: right;
        margin-left: 40px;

    }
</style>