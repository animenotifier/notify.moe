export interface Database {
	connect(): Promise<any>
	get(table: string, key: any)
	set(table: string, key: any, value: any)
	remove(table: string, key: any)
	forEach(table: string, func: (iterator: any) => any)
	filter(table: string, include: (iterator: any) => boolean)
	getMany(table: string, key: any[])
}