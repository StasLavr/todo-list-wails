export namespace task {
	
	export class Item {
	    id: number;
	    name?: string;
	    data?: string;
	    data_create?: string;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.data = source["data"];
	        this.data_create = source["data_create"];
	    }
	}

}

