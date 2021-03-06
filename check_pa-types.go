package publiccode

// checkPaTypes tells whether the pa-type is a valid type or not and returns it.
func (p *parser) checkPaTypes(key, paType string) (string, error) {
	for _, t := range paTypes {
		if t == paType {
			return paType, nil
		}
	}
	return paType, newErrorInvalidValue(key, "unknown pa-type: %s", paType)
}

// A paTypes represents a list of the types of public administration which is expected to use this software.
var paTypes = []string{
	// International.
	"cities",          // City
	"health-services", // Health service
	"police-forces",   // Police forces
	"schools",         // School of any level
	"universities",    // University of any level
	// Italy.
	"it-ag-turismo",    // Agenzie ed Enti per il Turismo
	"it-ag-lavoro",     // Agenzie ed Enti Regionali del Lavoro
	"it-ag-agricolo",   // Agenzie ed Enti Regionali di Sviluppo Agricolo
	"it-ag-formazione", // Agenzie ed Enti Regionali per la Formazione, la Ricerca e l'Ambiente
	"it-ag-fiscale",    // Agenzie Fiscali
	"it-ag-negoziale",  // Agenzie Regionali e Provinciale per la Rappresentanza Negoziale
	"it-ag-erogagric",  // Agenzie Regionali per le Erogazioni in Agricoltura
	"it-ag-sanita",     // Agenzie Regionali Sanitarie
	"it-ag-dirstudio",  // Agenzie, Enti e Consorzi Pubblici per il Diritto allo Studio Universitario
	"it-altrilocali",   // Altri Enti Locali
	"it-aci",           // Automobile Club Federati ACI
	"it-au-indip",      // Autorità Amministrative Indipendenti
	"it-au-ato",        // Autorità di Ambito Territoriale Ottimale
	"it-au-bacino",     // Autorità di Bacino
	"it-au-portuale",   // Autorità Portuali
	"it-az-edilizia",   // Aziende e Consorzi Pubblici Territoriali per l'Edilizia Residenziale
	"it-az-autonomo",   // Aziende ed Amministrazioni dello Stato ad Ordinamento Autonomo
	"hospital",         // Aziende Ospedaliere, Aziende Ospedaliere Universitarie, Policlinici e Istituti di Ricovero e Cura a Carattere Scientifico Pubblici
	"it-az-servizi",    // Aziende Pubbliche di Servizi alla Persona
	"it-az-sanita",     // Aziende Sanitarie Locali
	"it-camcom",        // Camere di Commercio, Industria, Artigianato e Agricoltura e loro Unioni Regionali
	"it-metro",         // Città Metropolitane
	"city",             // Comuni e loro Consorzi e Associazioni
	"it-montana",       // Comunità Montane e loro Consorzi e Associazioni
	"it-co-bacino",     // Consorzi di Bacino Imbrifero Montano
	"it-co-ricerca",    // Consorzi Interuniversitari di Ricerca
	"it-co-industria",  // Consorzi per l'Area di Sviluppo Industriale
	"it-co-locali",     // Consorzi tra Amministrazioni Locali
	"it-centrale",      // Presidenza del Consiglio dei Ministri, Ministeri e Avvocatura dello Stato
	"it-provincia",     // Province e loro Consorzi e Associazioni
	"police",           // Forze di Polizia ad Ordinamento Civile e Militare per la Tutela dell'Ordine e della Sicurezza Pubblica
	"it-regione",       // Regioni, Province Autonome e loro Consorzi e Associazioni
	"it-afam",          // Istituzioni per l'Alta Formazione Artistica, Musicale e Coreutica
	"school",           // Istituti di Istruzione Statale di Ogni Ordine e Grado
	"university",       // Università e Istituti di Istruzione Universitaria Pubblici
}
