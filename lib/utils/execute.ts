import * as bluebird from 'bluebird'
import { exec } from 'child_process'

const execAsync = bluebird.promisify((command: string, callback) => {
	exec(command, function(error, stdout, stderr) {
		callback(error, stdout)
	})
})

export function execute(command: string): bluebird<string> {
	return <bluebird<string>> execAsync(command)
}