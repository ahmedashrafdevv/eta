USE MCSTREE



CREATE TABLE StkTrEReceiptHead(
	Serial int IDENTITY(1,1) NOT NULL,
	receiptNumber int NOT NULL,
	storeCode int NOT NULL,
    deviceSerialNumber NVARCHAR(100),
	dateTimeIssued datetime NOT NULL,
	totalDiscountAmount real NOT NULL DEFAULT 0,
	totalAmount real NOT NULL,
	totalTax real NOT NULL,
	stkTr03Serial int NULL,
	accountSerial int NULL,
    paymentMethod NVARCHAR(100),
	deletedAt datetime NULL,
	created_at datetime NOT NULL DEFAULT GETDATE(),
	posted bit NOT NULL DEFAULT 0,
) ON [PRIMARY]



CREATE TABLE StkTrEReceiptDetails(
	Serial int IDENTITY(1,1) NOT NULL,
	headSerial int NOT NULL,
	itemSerial int NOT NULL,
	itemType varchar(10) NOT NULL DEFAULT ('EGS'),
	internalCode varchar(100) NOT NULL,
	itemCode varchar(100) NOT NULL,
	[description] TEXT NOT NULL,
	unitType varchar(10) NOT NULL DEFAULT('BOX'),
	unitPrice real NOT NULL,
	quantity int NOT NULL,
	netSale real NOT NULL,
	itemDiscount real NOT NULL DEFAULT 0,
	deletedAt datetime NULL,
	createdAt datetime NULL DEFAULT GETDATE(),
	deleted bit NULL
) ON [PRIMARY]

GO
CREATE PROCEDURE StkTr03ListByStoreConvertedDate (@converted SMALLINT = -1 , @storeCode SMALLINT = -1 , @fromDate VARCHAR(20) = null, @toDate VARCHAR(20) = null)
	AS
    BEGIN
        SELECT  o.Serial ,  o.DocNo ,  o.DocDate ,  ISNULL(o.Discount , 0) Discount  ,  o.TotalCash ,
			  ISNULL(o.SaleTax , 0) SaleTax    
	    FROM StkTr03 o 
		WHERE ISNULL(o.TotalCash , 0) != 0
		AND ISNULL(EtaConverted , 0) = CASE WHEN @Converted = -1 THEN ISNULL(EtaConverted , 0) ELSE  @Converted END 
		AND StoreCode = CASE WHEN @StoreCode = 0 THEN StoreCode ELSE @StoreCode END
		AND o.DocDate >= CASE WHEN @fromDate IS NULL THEN o.DocDate ELSE CONVERT(DATETIME, @fromDate, 102) END
		AND o.DocDate <= CASE WHEN @toDate IS NULL THEN o.DocDate ELSE CONVERT(DATETIME, @toDate, 102) END
    END



GO
ALTER PROCEDURE StkTr03ConvertEReceipt (@Serial INT)
	AS
    BEGIN
        DECLARE @receiptNumber INT
        DECLARE @storeCode int
        DECLARE @dateTimeIssued datetime
        DECLARE @totalDiscountAmount REAL
        DECLARE @totalAmount REAL
        DECLARE @totalTax REAL
        DECLARE @accountSerial int
		DECLARE @headSerial int
		


		
        SELECT  
                @storeCode = StoreCode , 
                @dateTimeIssued = DocDate , 
                @totalDiscountAmount = Discount , 
                @totalAmount = TotalCash , 
                @totalTax = ISNULL(SaleTax , 0)  , 
                @accountSerial = AccountSerial FROM StkTr03 WHERE Serial = @Serial

		SET @receiptNumber = (SELECT ISNULL(MAX(receiptNumber) , 1) FROM StkTrEReceiptHead WHERE storeCode = @storeCode)
        INSERT INTO StkTrEReceiptHead (receiptNumber , storeCode ,dateTimeIssued ,totalDiscountAmount ,totalAmount ,totalTax ,stkTr03Serial ,accountSerial , paymentMethod)
        VALUES (@receiptNumber ,@storeCode ,@dateTimeIssued ,@totalDiscountAmount ,@totalAmount ,@totalTax ,@Serial ,@accountSerial , 'CASH')

        SET @headSerial = SCOPE_IDENTITY();

        INSERT INTO StkTrEReceiptDetails ( headSerial , itemCode  , unitPrice , quantity , itemDiscount , itemSerial , internalCode , "description" )
        SELECT @headSerial , ISNULL(o.BarCodeUsed, '123') , o.Price , o.Qnt , ISNULL(o.Discount,0 ), o.ItemSerial ,  i.ItemCode , i.ItemName
        FROM    StkTr04 o JOIN StkMs01 i ON o.ItemSerial = i.Serial WHERE HeadSerial = @Serial  


		UPDATE StkTr03 SET EtaConverted = 1 WHERE "Serial" = @Serial
		SELECT SCOPE_IDENTITY() id
    END
