<script>
    import "@github/details-dialog-element";
    import Textfield from "@smui/textfield";
    import Icon from "@smui/textfield/icon";
    import HelperText from "@smui/textfield/helper-text";
    import List, {
        Item,
        Meta,
        Text,
        PrimaryText,
        SecondaryText,
    } from "@smui/list";
    import Card from "@smui/card";
    import { v1 as v1uuid } from "uuid";
    let accounts = [
        {
            id: "xxxxxxxxxx",
            address: "youre-team.1password.com",
            email: "email@example.com",
            secret: "A3-*****-****-****-****-****-****",
            password: "xxxxxxxxxxxxxxxxxxxxx",
            disabled: false,
        },
        {
            id: "vvvvvvvvvvvvvv",
            address: "youre-team.1password.com",
            email: "email2@example.com",
            secret: "A3-*****-****-****-****-****-****",
            password: "xxxxxxxxxxxxxxxxxxxxx",
        },
    ];
    let newAccountData = () => {
        return {
            id: null,
            address: "",
            email: "",
            secret: "",
            password: "",
        };
    };
    let data = newAccountData();
    let isEdit = false;
    let newAccount = () => {
        accounts = [
            ...accounts,
            {
                id: v1uuid(),
                address: data.address,
                email: data.email,
                secret: data.secret,
                password: data.password,
            },
        ];
        openEditAccount = false;
        console.log("newAccount", accounts);
        data = newAccountData();
    };
    let deleteAccountID = "";
    let deleteAccount = (id) => {
        console.log("deleteAccount ID:", id);
        accounts = accounts.filter((account) => account.id !== deleteAccountID);
    };
    let openDelete = false;
    let openSetting = false;
    let clicked = "Nothing yet.";

    let openEditAccount = false;
    let clickedInner = "Nothing yet.";
    let updateAccount = () => {
        isEdit = false;
        let accountDB = {
            address: data.address,
            email: data.email,
            secret: data.secret,
            password: data.password,
            id: data.id,
        };
        let objIndex = accounts.findIndex((obj) => obj.id == accountDB.id);
        console.log("Before update: ", accounts[objIndex]);
        accounts[objIndex] = accountDB;
        console.log("update: ", accounts[objIndex]);
        openEditAccount = false;
        data = newAccountData();
    };
    let editTitleMap = {
        true: "Edit",
        false: "Add",
    };
    let editButtonMap = {
        true: "Update",
        false: "Add",
    };
    let editApplyMap = {
        true: updateAccount,
        false: newAccount,
    };
    import Dialog, { Title, Content, Actions } from "@smui/dialog";
    import Button, { Label } from "@smui/button";
</script>

<Dialog
    bind:open={openSetting}
    aria-labelledby="simple-title"
    aria-describedby="simple-content"
>
    <!-- Title cannot contain leading whitespace due to mdc-typography-baseline-top() -->
    <Title id="simple-title">Account Setting</Title>
    <Content id="simple-content">
        <Card padded>
            <List class="demo-list" twoLine avatarList singleSelection>
                {#each accounts as account}
                    <Item disabled={account.disabled} nonInteractive>
                        <Text>
                            <PrimaryText>{account.address}</PrimaryText>
                            <SecondaryText>{account.email}</SecondaryText>
                        </Text>
                        <Meta>
                            <Button
                                on:click={() => {
                                    isEdit = true;
                                    data = account;
                                    openEditAccount = true;
                                }}
                            >
                                <Label>Edit</Label>
                            </Button>
                            <Button
                                on:click={() => {
                                    openDelete = true;
                                    deleteAccountID = account.id;
                                }}
                            >
                                <Label>Delete</Label>
                            </Button>
                        </Meta>
                    </Item>
                {/each}
            </List>
        </Card>
        <Button
            on:click={() => {
                isEdit = false;
                data = newAccountData();
                openEditAccount = true;
            }}
        >
            <Label>Add Account</Label>
        </Button>
    </Content>
    <Actions>
        <Button on:click={() => (clicked = "cancel")}>
            <Label>Cancel</Label>
        </Button>
        <Button on:click={() => (clicked = "apply")}>
            <Label>Apply</Label>
        </Button>
    </Actions>
</Dialog>
<Dialog
    bind:open={openEditAccount}
    aria-labelledby="dialog-inner"
    aria-describedby="dialog-inner-content"
>
    <Title id="dialog-inner">{editTitleMap[isEdit]} account</Title>
    <Content id="dialog-inner-content">
        <Textfield
            bind:value={data.address}
            label="Signin address"
            style="min-width: 350px;"
        >
            <Icon class="material-icons" slot="leadingIcon">home</Icon>
            <HelperText slot="helper">youre-team.1password.com</HelperText>
        </Textfield>
        <Textfield
            bind:value={data.email}
            label="Email"
            style="min-width: 350px;"
        >
            <Icon class="material-icons" slot="leadingIcon">email</Icon>
            <HelperText slot="helper">email@example.com</HelperText>
        </Textfield>
        <Textfield
            bind:value={data.secret}
            type="password"
            label="Secret key"
            style="min-width: 350px;"
        >
            <Icon class="material-icons" slot="leadingIcon">key</Icon>
            <HelperText slot="helper"
                >A3-*****-****-****-****-****-****</HelperText
            >
        </Textfield>
        <Textfield
            bind:value={data.password}
            type="password"
            label="Master password"
            style="min-width: 350px;"
        >
            <Icon class="material-icons" slot="leadingIcon">password</Icon>
        </Textfield>
    </Content>
    <Actions>
        <Button on:click={() => (clickedInner = "cancel")}>
            <Label>Cancel</Label>
        </Button>
        <Button on:click={() => editApplyMap[isEdit]()}>
            <Label>{editButtonMap[isEdit]}</Label>
        </Button>
    </Actions>
</Dialog>

<Dialog
    bind:open={openDelete}
    aria-labelledby="delete-title"
    aria-describedby="delete-content"
    on:SMUIDialog:closed={(e) => {
        e.detail.action === "yes" && deleteAccount();
    }}
>
    <!-- Title cannot contain leading whitespace due to mdc-typography-baseline-top() -->
    <Title id="delete-title">Delete Account</Title>
    <Content id="delete-content">Do you really want to delete this?</Content>
    <Actions>
        <Button action="no" default>
            <Label>No</Label>
        </Button>
        <Button action="yes">
            <Label>Yes</Label>
        </Button>
    </Actions>
</Dialog>

<Button on:click={() => (openSetting = true)}>Setting</Button>

<pre class="status">Clicked: {clicked}</pre>
