import React from "react";

import { Box, Flex, Heading, Text, Stack } from "@chakra-ui/react";
import { LocalePicker } from "components";
import { intl } from "locale";
import { FaTimes } from "react-icons/fa";

function Unhealthy() {
	return (
		<>
			<Stack h={10} m={4} justify={"end"} direction={"row"}>
				<LocalePicker className="text-right" />
			</Stack>
			<Box textAlign="center" py={10} px={6}>
				<Box display="inline-block">
					<Flex
						flexDirection="column"
						justifyContent="center"
						alignItems="center"
						bg={"red.500"}
						rounded={"50px"}
						w={"55px"}
						h={"55px"}
						textAlign="center">
						<FaTimes size={"30px"} color={"white"} />
					</Flex>
				</Box>
				<Heading as="h2" size="xl" mt={6} mb={2}>
					{intl.formatMessage({
						id: "unhealthy.title",
						defaultMessage: "Nginx Proxy Manager is unhealthy",
					})}
				</Heading>
				<Text color={"gray.500"}>
					{intl.formatMessage({
						id: "unhealthy.body",
						defaultMessage:
							"We'll continue to check the health and hope to be back up and running soon!",
					})}
				</Text>
			</Box>
		</>
	);
}

export { Unhealthy };
