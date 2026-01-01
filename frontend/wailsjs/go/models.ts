export namespace domain {
	
	export class EnvStatus {
	    Name: string;
	    Scope: string;
	    Exists: boolean;
	    Correct: boolean;
	    Missing: string[];
	    Value: string;
	
	    static createFrom(source: any = {}) {
	        return new EnvStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Scope = source["Scope"];
	        this.Exists = source["Exists"];
	        this.Correct = source["Correct"];
	        this.Missing = source["Missing"];
	        this.Value = source["Value"];
	    }
	}
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

