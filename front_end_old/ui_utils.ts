export function formatDate(date: Date) {
    console.log(date)
    // get the type of the date
    if (typeof date === 'string') {
        date = new Date(date);
    }

    return date.toLocaleDateString('en-US', {
        month: 'short', // "short" for abbreviated month name
        day: 'numeric', // numeric day
        year: 'numeric', // numeric year
    });
}