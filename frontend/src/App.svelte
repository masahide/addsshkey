<script>
    import Accounts from "./Accounts.svelte";
    import LoadedKeys from "./loadedkeys.svelte";
    import Keys from "./Keys.svelte";
    import Card from "@smui/card";
    import Paper, { Title, Content } from "@smui/paper";
    import Button, { Label, Icon } from "@smui/button";

    let errMsg = "";

    import Tab from "@smui/tab";
    import TabBar from "@smui/tab-bar";

    let tabs = [
        {
            icon: "key",
            label: "Keys",
        },
        {
            icon: "support_agent",
            label: "Loaded keys",
        },
        {
            icon: "settings",
            label: "Settings",
        },
    ];
    let active = tabs[0];
</script>

<main>
    <div id="input" data-wails-no-drag>
        {#if errMsg}
            <div class="paper-container">
                <Paper color="primary">
                    <Title>Primary Paper</Title>
                    <Content>
                        {errMsg}
                    </Content>
                </Paper>
            </div>
            <div class="flash flash-error">
                <button
                    class="flash-close js-flash-close"
                    type="button"
                    aria-label="Close"
                    on:click={() => {
                        errMsg = "";
                    }}
                />
            </div>
        {/if}
        <div>
            <TabBar {tabs} let:tab bind:active>
                <Tab {tab}>
                    <Icon class="material-icons">{tab.icon}</Icon>
                    <Label>{tab.label}</Label>
                </Tab>
            </TabBar>
        </div>
        {#if active.label === "Settings"}
            <Card><Accounts /></Card>
        {:else if active.label === "Loaded keys"}
            <Card><LoadedKeys /></Card>
        {:else if active.label === "Keys"}
            <Keys />
        {/if}
    </div>
</main>

<style>
    #input {
        width: 95%;
        margin: 5% auto;
    }

    button {
        -webkit-appearance: default-button;
        padding: 6px;
    }
</style>
