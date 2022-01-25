<script>
    import Card from "@smui/card";
    import Checkbox from "@smui/checkbox";
    import Autocomplete from "@smui-extra/autocomplete";
    import Chip, { Set, TrailingAction, Text } from "@smui/chips";
    import Button, { Label } from "@smui/button";
    import { v1 as v1uuid } from "uuid";
    import CircularProgress from "@smui/circular-progress";

    // tags
    let tags = [];

    function addTagsChip() {
        console.log("value:",valueStandard);
        if (valueStandard) {
            tags.push({ id: v1uuid(), v: valueStandard });
            tags = tags;
            valueStandard = "";
        }
    }

    let fruits = [
        "Apple",
        "Orange",
        "Banana",
        "Mango",
        "Lemon",
        "Cherry",
        "Blueberry",
        "Grape",
        "Strawberry",
    ];
    let valueStandard = "";

    let counter = 0;

    async function searchItems(input) {
        if (input === "") {
            return [];
        }
        valueStandard = input;

        // Pretend to have some sort of canceling mechanism.
        const myCounter = ++counter;

        // Pretend to be loading something...
        await new Promise((resolve) => setTimeout(resolve, 100));

        if (myCounter !== counter) {
            // This means the function was called again, so we should cancel.
            // This is how you tell Autocomplete to cancel this search. It won't
            // replace the results of any subsequent search that has already finished.
            return false;
        }

        // Return a list of matches.
        return fruits.filter((item) =>
            item.toLowerCase().includes(input.toLowerCase())
        );
    }

    let data = [
        {
            id: 1,
            type: "Tobie",
            finger_print: "Vint",
            comment: "tvint0@fotki.com",
        },
        {
            id: 2,
            type: "Zacharias",
            finger_print: "Cerman",
            comment: "zcerman1@sciencedirect.com",
        },
        {
            id: 3,
            type: "Gerianna",
            finger_print: "Bunn",
            comment: "gbunn2@foxnews.com",
        },
        {
            id: 4,
            type: "Bee",
            finger_print: "Saurin",
            comment: "bsaurin3@live.com",
        },
        {
            id: 5,
            type: "Meyer",
            finger_print: "Garnul",
            comment: "mgarnul4@yellowbook.com",
        },
        {
            id: 6,
            type: "Frederich",
            finger_print: "Benley",
            comment: "fbenley5@ameblo.jp",
        },
        {
            id: 7,
            type: "Becki",
            finger_print: "Criag",
            comment: "bcriag6@washingtonpost.com",
        },
        {
            id: 8,
            type: "Nichols",
            finger_print: "Risom",
            comment: "nrisom7@google.com.br",
        },
        {
            id: 9,
            type: "Ron",
            finger_print: "Menendes",
            comment: "rmenendes8@prnewswire.com",
        },
        {
            id: 10,
            type: "Thane",
            finger_print: "Gammill",
            comment: "tgammill9@com.com",
        },
        {
            id: 11,
            type: "Ramonda",
            finger_print: "Yakobowitch",
            comment: "ryakobowitcha@hibu.com",
        },
        {
            id: 12,
            type: "Rosalynd",
            finger_print: "Connechie",
            comment: "rconnechieb@chicagotribune.com",
        },
        {
            id: 13,
            type: "Vinny",
            finger_print: "Hissett",
            comment: "vhissettc@4shared.com",
        },
        {
            id: 14,
            type: "Mylo",
            finger_print: "Tarbath",
            comment: "mtarbathd@behance.net",
        },
        {
            id: 15,
            type: "Pierson",
            finger_print: "Linny",
            comment: "plinnye@altervista.org",
        },
        {
            id: 16,
            type: "Adi",
            finger_print: "Bassano",
            comment: "abassanof@about.com",
        },
        {
            id: 17,
            type: "Paxon",
            finger_print: "Claypole",
            comment: "pclaypoleg@storify.com",
        },
        {
            id: 18,
            type: "Marcella",
            finger_print: "Mogey",
            comment: "mmogeyh@weebly.com",
        },
        {
            id: 19,
            type: "Fiann",
            finger_print: "Winsborrow",
            comment: "fwinsborrowi@icq.com",
        },
        {
            id: 20,
            type: "Moria",
            finger_print: "Druhan",
            comment: "mdruhanj@friendfeed.com",
        },
        {
            id: 21,
            type: "Ransom",
            finger_print: "Beidebeke",
            comment: "rbeidebekek@webnode.com",
        },
        {
            id: 22,
            type: "Ev",
            finger_print: "Scullion",
            comment: "escullionl@opensource.org",
        },
        {
            id: 23,
            type: "Etta",
            finger_print: "Rockwell",
            comment: "erockwellm@alexa.com",
        },
        {
            id: 24,
            type: "Rorke",
            finger_print: "Daft",
            comment: "rdaftn@list-manage.com",
        },
        {
            id: 25,
            type: "Waly",
            finger_print: "Kern",
            comment: "wkerno@wix.com",
        },
        {
            id: 26,
            type: "Wendell",
            finger_print: "McKyrrelly",
            comment: "wmckyrrellyp@nydailynews.com",
        },
        {
            id: 27,
            type: "Steward",
            finger_print: "Dillet",
            comment: "sdilletq@artisteer.com",
        },
        {
            id: 28,
            type: "Catherina",
            finger_print: "Cestard",
            comment: "ccestardr@msn.com",
        },
        {
            id: 29,
            type: "Colin",
            finger_print: "Large",
            comment: "clarges@posterous.com",
        },
        {
            id: 30,
            type: "Joann",
            finger_print: "Meineking",
            comment: "jmeinekingt@vkontakte.ru",
        },
    ];
    function loadKeys() {
        let res = data.map((row) => {
            return {
                selected: false,
                id: row.id,
                type: row.type,
                comment: row.comment,
            };
        });
        return res;
    }
    function clickrow(id) {
        console.log("click " + id);
        list.forEach((row, i) => {
            if (row.id === id) {
                list[i].selected = !list[i].selected;
            }
        });
    }

    let list = loadKeys();

</script>

<Card padded>
    <div>
        <Autocomplete
            search={searchItems}
            showMenuWithNoInput={false}
            label="Tags"
        >
            <Text
                slot="loading"
                style="display: flex; width: 100%; justify-content: center; align-items: center;"
            >
                <CircularProgress
                    style="height: 24px; width: 24px;"
                    indeterminate
                />
            </Text>
        </Autocomplete>
        <Button on:click={addTagsChip}><Label>Add</Label></Button>
        <Set chips={tags} let:chip key={(chip) => chip.id} input>
            <Chip {chip}>
                <Text>{chip.v}</Text>
                <TrailingAction icon$class="material-icons"
                    >cancel</TrailingAction
                >
            </Chip>
        </Set>
    </div>
</Card>
<Card padded>
    <div class="mdc-data-table">
        <table class="mdc-data-table__table">
            <thead>
                <tr class="mdc-data-table__header-row">
                    <th class="mdc-data-table__header-cell" />
                    <th class="mdc-data-table__header-cell">番号</th>
                    <th class="mdc-data-table__header-cell">KeyType</th>
                    <th class="mdc-data-table__header-cell"
                        >FingerPrintSHA256</th
                    >
                    <th class="mdc-data-table__header-cell">Comment</th>
                </tr>
            </thead>
            <tbody class="mdc-data-table__content">
                {#each list as row}
                    <tr class="mdc-data-table__row" on:click={clickrow(row.id)}>
                        <td class="mdc-data-table__cell"
                            ><Checkbox bind:checked={row.selected} /></td
                        >
                        <td class="mdc-data-table__cell">{row.id}</td>
                        <td class="mdc-data-table__cell">{row.type}</td>
                        <td class="mdc-data-table__cell">{row.finger_print}</td>
                        <td class="mdc-data-table__cell">{row.comment}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
</Card>

<style>
    th,
    td {
        vertical-align: middle;
    }
</style>
