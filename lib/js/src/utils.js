function is_null(value) {
    return value === null || value === undefined;
}

class Logger {
    static log(message) {
        console.log(message);
    }

    static error(message) {
        console.error(message);
    }
}
