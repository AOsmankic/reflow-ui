<script>
    import Logo from './Logo.svelte'
    import Hamburger from './Hamburger.svelte'

    export let sidebar = false

</script>

<header class="flex justify-between bg-black-200 p-2 items-center text-black-600">
    <nav class="flex">
        <Hamburger bind:open={sidebar}/>
    </nav>

</header>
