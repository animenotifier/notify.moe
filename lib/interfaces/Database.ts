export interface Database {
	connect(): Promise<any>
	get(table: string, key: any): Promise<any>
	set(table: string, key: any, value: any): Promise<void>
	remove(table: string, key: any): Promise<void>
	all(table: string): Promise<Array<any>>
	forEach(table: string, func: (iterator: any) => void): Promise<void>
	filter(table: string, include: (iterator: any) => boolean): Promise<any>
	getMany(table: string, key: any[]): Promise<Array<any>>
}