USE ETA_Master



CREATE TABLE EtaLog(
	Serial int IDENTITY(1,1) NOT NULL,
	internalID VARCHAR(100) NOT NULL,
	submissionID VARCHAR(100),
	storeCode int NOT NULL,
    serials TEXT NOT NULL,
    logText TEXT NOT NULL,
	posted bit NOT NULL DEFAULT 0,
	created_at datetime NOT NULL DEFAULT GETDATE(),
) ON [PRIMARY]

GO

CREATE PROC EtaLogInset(
    @internalID VARCHAR(100),
    @submissionID VARCHAR(100),
    @storeCode int,
    @serials TEXT,
    @logText TEXT,
    @posted bit

)
AS 
BEGIN
    INSERT INTO EtaLog (
        internalID,
        submissionID,
        storeCode,
        serials,
        logText,
        posted
    ) VALUES (
        @internalID,
        @submissionID,
        @storeCode,
        @serials,
        @logText,
        @posted
    )
END
