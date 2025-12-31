export namespace domain {
	
	export class PathStatus {
	    path: string;
	    exists: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PathStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.exists = source["exists"];
	    }
	}

}

