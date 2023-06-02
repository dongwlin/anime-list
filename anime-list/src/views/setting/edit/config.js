export const config = {
    tableHeader: [
        {
            label: 'ID',
            prop: 'id'
        },
        {
            label: 'Name',
            prop: 'name',
        },
        {
            label: 'Type',
            prop: 'type'
        },
        {
            label: 'Day',
            prop: 'day'
        },
        {
            label: 'Status',
            prop: 'status'
        },
        {
            label: 'Operations',
            prop: 'operations'
        }
    ],
    type: {
        none: -1,
        network: 0,
        local: 1
    },
    day: {
        none: -1,
        Sunday: 0,
        Monday: 1,
        Tuesday: 2,
        Wednesday: 3,
        Thursday: 4,
        Friday: 5,
        Saturday: 6
    },
    createFormRules: {
        name: [
            {
                required: true,
                message: 'Please input name',
                trigger: 'blur'
            },
        ],
        status: [
            {
                required: true,
                message: 'Please select status',
                trigger: 'change',
            },
        ],
        type: [
            {
                required: true,
                message: 'Please select type',
                trigger: 'change',
            },
        ],
        dir: [
            {
                required: false,
                message: 'Please input dir',
                trigger: 'blur',
            },
        ],
        url: [
            {
                required: false,
                message: 'Please input url',
                trigger: 'blur',
            },
        ],
        img: [
            {
                required: false,
                message: 'Please input img',
                trigger: 'blur',
            },
        ],
        day: [
            {
                required: true,
                message: 'Please select day',
                trigger: 'change'
            },
        ],
    },
    updateFormRules: {
        name: [
            {
                required: true,
                message: 'Please input name',
                trigger: 'blur'
            },
        ],
        status: [
            {
                required: true,
                message: 'Please select status',
                trigger: 'change',
            },
        ],
        type: [
            {
                required: true,
                message: 'Please select type',
                trigger: 'change',
            },
        ],
        dir: [
            {
                required: false,
                message: 'Please input dir',
                trigger: 'blur',
            },
        ],
        url: [
            {
                required: false,
                message: 'Please input url',
                trigger: 'blur',
            },
        ],
        img: [
            {
                required: false,
                message: 'Please input img',
                trigger: 'blur',
            },
        ],
        day: [
            {
                required: true,
                message: 'Please select day',
                trigger: 'change'
            },
        ],
    }
};
