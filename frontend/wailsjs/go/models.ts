export namespace main {
	
	export class ResultadoNReinas {
	    tablero: string[];
	    logs: string[];
	
	    static createFrom(source: any = {}) {
	        return new ResultadoNReinas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tablero = source["tablero"];
	        this.logs = source["logs"];
	    }
	}

}

