export namespace domain {
	
	export class EnvStatus {
	    name: string;
	    scope: string;
	    exists: boolean;
	    correct: boolean;
	    missing: string[];
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new EnvStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.scope = source["scope"];
	        this.exists = source["exists"];
	        this.correct = source["correct"];
	        this.missing = source["missing"];
	        this.value = source["value"];
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
	export class ShortcutStatus {
	    name: string;
	    source: string;
	    target: string;
	    exists: boolean;
	    correct: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ShortcutStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.source = source["source"];
	        this.target = source["target"];
	        this.exists = source["exists"];
	        this.correct = source["correct"];
	    }
	}
	export class SoftwareStatus {
	    name: string;
	    rootDir: string;
	    exists: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SoftwareStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.rootDir = source["rootDir"];
	        this.exists = source["exists"];
	    }
	}

}

