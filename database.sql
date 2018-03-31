CREATE TABLE IF NOT EXISTS map
(
	id INT PRIMARY KEY AUTO_INCREMENT,
	data TEXT NOT NULL,
	published SMALLINT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS score
(
	id INT PRIMARY KEY AUTO_INCREMENT,
	mapId INT NOT NULL,
	name VARCHAR(100),
	moves INT
	/* entryDate TIMESTAMP NOT NULL DEFAULT NOW() */
	/* for simplicity sake, mapId is not declared as a foreign key to the map table */
);

DROP VIEW IF EXISTS vwScoresForMap;
CREATE VIEW vwScoresForMap AS
	SELECT id, mapId, name, moves FROM score ORDER BY moves ASC, id ASC;

DELIMITER $$
DROP FUNCTION IF EXISTS funGetScoreIdPosition$$

CREATE FUNCTION funGetScoreIdPosition(scoreId INT) RETURNS INT DETERMINISTIC
BEGIN
  SELECT 0 INTO @rowNumber;

  SELECT rowNumber INTO @rtn FROM
    (SELECT (@rowNumber := @rowNumber + 1), @rowNumber AS rowNumber, id FROM score WHERE mapId = (SELECT mapId FROM score WHERE id=scoreId) ORDER BY moves ASC, id ASC) AS t
  WHERE id = scoreId;

  RETURN @rtn;
END $$

DELIMITER ;