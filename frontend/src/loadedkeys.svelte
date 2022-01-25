<script>
    import { dndzone } from "svelte-dnd-action";
    import Paper, { Title, Subtitle, Content } from "@smui/paper";
    import LayoutGrid, { Cell } from "@smui/layout-grid";
    import Button, { Label } from "@smui/button";

    function handleSort(e) {
        items = e.detail.items;
    }
    let items = [
        {
            id: "333333333",
            type: "I",
            fingerprint: "hogege",
            comment: "aaaaaaa",
        },
        { id: "222", type: "I", fingerprint: "fffff", comment: "aaaaaaa" },
        {
            id: "1111111",
            type: "I",
            fingerprint: "bbbbbbbbb",
            comment: "aaaaaaa",
        },
    ];
    function deleteKey() {
        console.log("delete");
    }
</script>

<section
    use:dndzone={{ items }}
    on:consider={handleSort}
    on:finalize={handleSort}
>
    {#each items as item (item.id)}
        <div>
            <LayoutGrid>
                <Cell span={2}>{item.type}</Cell>
                <Cell span={8}
                    ><Cell span={12}>{item.fingerprint}</Cell>
                    <Cell span={12}>{item.comment}</Cell>
                </Cell>
<Cell span={2}>
        <Button on:click={deleteKey}><Label>Delete</Label></Button>
</Cell>
            </LayoutGrid>
        </div>
    {/each}
</section>

<style>
    div {
        border: 1px solid black;
        margin: 0.2em;
        padding: 0.3em;
    }
</style>
