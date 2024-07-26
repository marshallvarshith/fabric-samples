'use strict';

const { Contract } = require('fabric-contract-api');
const { exec } = require('child_process');

class SmartLockContract extends Contract {
    async lock(ctx) {
        exec('python3 smart_lock.py lock', (error, stdout, stderr) => {
            if (error) {
                console.error(`exec error: ${error}`);
                return;
            }
            console.log(`stdout: ${stdout}`);
            console.error(`stderr: ${stderr}`);
        });
    }

    async unlock(ctx) {
        exec('python3 smart_lock.py unlock', (error, stdout, stderr) => {
            if (error) {
                console.error(`exec error: ${error}`);
                return;
            }
            console.log(`stdout: ${stdout}`);
            console.error(`stderr: ${stderr}`);
        });
    }
}

module.exports = SmartLockContract;
