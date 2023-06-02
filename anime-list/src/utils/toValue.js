import {config} from "@/views/setting/edit/config.js";

export const toType = (type_value) => {
    switch (type_value) {
        case config.type.none:
            return 'none';
        case config.type.network:
            return 'network';
        case config.type.local:
            return 'local';
        default:
            return 'none';
    }
}

export const toDay = (day_value) => {
    switch (day_value) {
        case config.day.none:
            return 'none';
        case config.day.Sunday:
            return 'Sunday';
        case config.day.Monday:
            return 'Monday';
        case config.day.Tuesday:
            return 'Tuesday';
        case config.day.Wednesday:
            return 'Wednesday';
        case config.day.Thursday:
            return 'Thursday';
        case config.day.Friday:
            return 'Friday';
        case config.day.Saturday:
            return 'Saturday';
        default:
            return 'none';
    }
}
