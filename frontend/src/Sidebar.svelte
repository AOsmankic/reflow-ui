<script>
    import { Pane, Splitpanes } from 'svelte-splitpanes';
    import Profile from "./Profile.svelte";
    import PreviewProfile from "./PreviewProfile.svelte";

    import {
        GetAllProfiles
    } from '../wailsjs/go/main/App.js'
    import {onMount} from "svelte";
    export let open = false
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
    })
    // $: if (activeProfile !== oldProfile){
    //     oldProfile = activeProfile
    //     if (activeProfile != null) {
    //         previewPane.updatePreview()
    //     }
    // }
</script>
<aside class="absolute w-full h-full bg-gray-200 border-r-2 shadow-lg" class:open>
    <Splitpanes class="default-theme" style="height: 100%">
        <Pane minSize={30} size={30}>
            <Profile {profiles} bind:activeProfile={activeProfile}/>
        </Pane>
        <Pane>
            <PreviewProfile {activeProfile} bind:this={previewPane}/>
        </Pane>
    </Splitpanes>
</aside>

<style>
    aside {
        left: -100%;
        transition: left 0.3s ease-in-out
    }

    .open {
        left: 0
    }
</style>