const machine = createMachine({
    initialState:'start',
    start:{
        actions:{
            openBracket(){
                console.log('Start state: open bracket');
            },
        },
        transitions:{
            '[': {target:'title', action: () => console.log('Transition to title state')},
            '\n': {target:'end', action: () => console.log('Transition to end state')}
        },
    },
    title:{
        actions:{
            words(char){
                console.log(`Title state adding char ${char}`)
            },
            closeBracket(){
                console.log('Title state close bracket')
            },
        },
        transitions:{
            ']':{target:'last',action:()=>console.log('Transition title state to last state')},
        },
    },
    last:{
        actions:{
            words(char){
                console.log(`Last state adding char ${char}`);
            },
            closeBracket(){
                console.log('Last state close bracket')
            },
        },
        transitions:{
            'end':{}
        }
    }
})