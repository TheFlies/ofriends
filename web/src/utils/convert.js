import moment from "moment"

export function getHumanDate(timestamp) {
    return moment(timestamp).format('MMMM Do YYYY, h:mm a');
}